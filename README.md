# bday-email
Python/Go script that will send a personalized birthday message over email

## Python Version
* Uses post request to my [Go Back-End](https://github.com/go-backend) server to complete the emailing process
* Rest of the code is done in the Python script..
   * Figures out if anyone in my Database has a birthday today and if so will send a personalized email with random joke

## Go Version
* No need to use the back-end server
* All code necessary is in this repo
* Main.go imports our emailer package
* **I still haven't figured out a scheduler with go to run the script each morning**

## Random Joke
This is a random joke pulled from API by [KegenGuyII](https://github.com/KegenGuyll/DadJokes_API)
* Could add in where when joke is sent to person we append a set with the joke number so we don't retell any jokes to people

## Secure
* Decided to add a layer of security to my back-end emailer
* Since this is open-source, anyone could see the url for my back-end api
    * Knowing this, there is now required authentification when trying to email a birthday message
* BASE_URL is also now hidden as well

# To Do (Completed)
* Recreate this code in Go and remove the api calls to my back-end
* I can just have the Go script run as a stand-alone application
* This will be good for learning go and I think a better implementation for this script

### Above To Dos are completed, however need to figure out a run scheduler
