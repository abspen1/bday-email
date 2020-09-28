import os
import redis

def get_bdays():
    client = redis.Redis(host="10.10.10.1", port=6379, db=6,
                         password=os.getenv("REDIS_PASS"))
    dictionary = client.hgetall('bday')
    for item, key in dictionary.items():
        print(item, key, "\n\n")

def remove_bdays():
    client = redis.Redis(host="10.10.10.1", port=6379, db=6,
                         password=os.getenv("REDIS_PASS"))
    client.hdel("bday", "First Last Mon 01")
    
def main(name, birthdate, email):
    client = redis.Redis(host="10.10.10.1", port=6379, db=6,
                         password=os.getenv("REDIS_PASS"))
    # Using a hash
    # client.hset('bday', 'First Last Mon 01', 'email@domain')
    concatenate = name + " " + birthdate
    client.hset('bday', concatenate, email)
    

if __name__ == "__main__":
    remove_bdays()
    get_bdays()
