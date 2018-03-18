from random import randint

global currentState
def creditsState():
    print("This game is developed by Preslav Mihaylov")
    input()
    return mainMenuState

def playState():
    print("What is your bet?")

    print("1 - heads")
    print("2 - tails")

    userInput = input()
    if userInput != "1" and userInput != "2":
        print("Invalid input. Try again!")
        return currentState

    result = randint(1, 2)
    guess = int(userInput)

    if guess == result:
        print("You got it right!")
    elif result == 1:
        print("Wrong! It was heads")
        input()
        return mainMenuState
    else:
        print("Wrong! It was tails")
        input()
        return mainMenuState

    return currentState

def mainMenuState():
    print("Welcome to the coin game! What would you like to do?")
    print("1 - play")
    print("2 - see credits section")

    userInput = input()
    if userInput == "1":
        return playState
    elif userInput == "2":
        return creditsState
    else:
        print("Invalid option. Try again!")

currentState = mainMenuState
while True:
    currentState = currentState()
