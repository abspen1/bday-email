import os
import redis

def main(name, birthdate, email):
    client = redis.Redis(host="10.10.10.1", port=6379, db=6,
                         password=os.getenv("REDIS_PASS"))
    # Using a hash
    # client.hset('bday', 'First Last Mon 01', 'email@domain')
    concatenate = name + " " + birthdate
    client.hset('bday', concatenate, email)

if __name__ == "__main__":
    main()
