import datetime
import json
import os
import random
import redis
import requests
import schedule
import time

BASE_URL = os.getenv("BASE_URL")

JOKE_URL = "https://us-central1-dadsofunny.cloudfunctions.net/DadJokes/random/type/"


def add_bday(name, birthdate, email):
    client = redis.Redis(host=os.getenv("REDIS_HOST"), port=6379, db=6,
                         password=os.getenv("REDIS_PASS"))
    # Using a hash
    # client.hset('bday', 'First Last Mon 01', 'email@domain')
    concatenate = name + " " + birthdate
    client.hset('bday', concatenate, email)


def get_bdays():
    client = redis.Redis(host=os.getenv("REDIS_HOST"), port=6379, db=6,
                         password=os.getenv("REDIS_PASS"))
    dictionary = client.hgetall('bday')
    for item, key in dictionary.items():
        print(item, key, "\n\n")


def get_joke():
    types = ("general", "knock-knock")
    index = random.randint(0,1)
    url = JOKE_URL + types[index]
    r = requests.get(url)
    return json.loads(r.content)


def get_formatted_date():
    # Get date in format 'Mon 00'
    date = time.ctime()[4:10]
    return date


def send_email(name, email, auth):
    joke = get_joke()
    for item in joke:
        setup = item["setup"]
        punchline = item["punchline"]
    dictionary = {"Name": name, "Email": email, "JokeSetup": setup, "JokePunchLine": punchline, "Auth": auth}
    dictionary = json.dumps(dictionary, indent=4)
    x = requests.post(BASE_URL, data=dictionary)

    print(x.text)


def main():
    date = get_formatted_date()
    client = redis.Redis(host=os.getenv("REDIS_HOST"), port=6379, db=6,
                         password=os.getenv("REDIS_PASS"))
    birthday_dict = client.hgetall('bday')

    for name_date, email in birthday_dict.items():
        if date == name_date.decode("utf-8")[-6:]:
            name = name_date.decode("utf-8")[:-7]
            email = email.decode("utf-8")
            auth = os.getenv("BACK-END-AUTH")
            send_email(name, email, auth)


add_bday("TEST", "Dec 20", "abspencer2097@yahoo.com")
get_bdays()
schedule.every().day.at("06:00").do(main)

while True:
    try:
        schedule.run_pending()
        time.sleep(1)
    except Exception as identifier:
        print(identifier)
        time.sleep(1)
