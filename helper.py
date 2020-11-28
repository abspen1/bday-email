import os
import redis

client = redis.Redis(host="10.10.10.1", port=6379, db=6,
                     password=os.getenv("REDIS_PASS"))

def get_bdays():
    dictionary = client.hgetall('bday')
    for item, key in dictionary.items():
        print(item, key, "\n\n")

def remove_bdays(name, birthdate):
    concatenate = name + " " + birthdate
    client.hdel("bday", concatenate)
    
def main(name, birthdate, email):
    # Using a hash
    # client.hset('bday', 'First Last Mon 01', 'email@domain')
    concatenate = name + " " + birthdate
    client.hset('bday', concatenate, email)
    

if __name__ == "__main__":
    # remove_bdays()
    print(client.get('back-end-auth'))
