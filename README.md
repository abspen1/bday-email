# bday-email
Python script that will send a personalized birthday message over email

## Random Joke
This is a random joke pulled from API by [KegenGuyII](https://github.com/KegenGuyll/DadJokes_API)
* Could add in where when joke is sent to person we append a set with the joke number so we don't retell any jokes to people

## Main
Pretty simple code
* Get our dictionary of people's name/birthday(concatenated) and their email
* Check if any of them have a birthday today
* If they do, post request our backend email sender on my website

## Secure
* Decided to add a layer of security to my back-end emailer
* Since this is open-source, anyone could see the url for my back-end api
    * Knowing this, there is now required authentification when trying to email a birthday message
* BASE_URL is also now hidden as well

# To Do
* Recreate this code in Go and remove the api calls to my back-end
* I can just have the Go script run as a stand-alone application
* This will be good for learning go and I think a better implementation for this script