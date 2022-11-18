import pyrebase

firebaseConfig = {
  "apiKey": "AIzaSyA4JKLYbruafRG3PkA_if8dISYsHVSSEGM",
  "authDomain": "studentmgmtsystem-ab34f.firebaseapp.com",
  "projectId": "studentmgmtsystem-ab34f",
  "storageBucket": "studentmgmtsystem-ab34f.appspot.com",
  "messagingSenderId": "72346977656",
  "appId": "1:72346977656:web:7967ccdde992407ce8f48d",
  "databaseURL": ""
}

firebase=pyrebase.initialize_app(firebaseConfig)
auth=firebase.auth()

#Login function
def login():
    print("Log in...")
    email=input("Enter email: ")
    password=input("Enter password: ")
    try:
        login = auth.sign_in_with_email_and_password(email, password)
        print("Successfully logged in!")
        print("refresh token", login['idToken'])
    except:
        print("Invalid email or password")
    return

def signup():
    print("Sign up...")
    email = input("Enter email: ")
    password=input("Enter password: ")
    try:
        user = auth.create_user_with_email_and_password(email, password)
        ask=input("Do you want to login?[y/n]")
        if ask=='y':
            login()
    except: 
        print("Email already exists")
    return

#Main
ans=input("Are you a new user?[y/n]")

if ans == 'n':
    login()
elif ans == 'y':
    signup()