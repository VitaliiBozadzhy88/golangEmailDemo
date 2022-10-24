# 📨 📨 📨 Email 

### Project description

#### A simple microservice, where you can add email to DB, confirm it, control the process of confirming email.

### Features:

* ✅ add user`s email
* ✅ control activation code
* ✅ control execution status

### 📝 List of technologies used in the project

* ✅ Intellij IDEA
* ✅ MySQL
* ✅ GoLang

### 📀 Preparing to launch a project

* ✅ Clone project (need to push Code button and choose in what way you want to get project)
* ✅ Open with IntelliJ IDEA (if you do not have this program - download it [here](https://www.jetbrains.com/idea/download/#section=mac))
* ✅ You will also need MySQL to this project. Setup instructions [here](https://www.youtube.com/watch?v=xaPuXh8IFIU) and [here](https://www.youtube.com/watch?v=ImqxBiv5yIY)
* ✅ So, we have Intellij, Browser(Safari, Google Chrome etc.) and MySql
* ✅ Find in [User](model/User.go) fields - USER, PASS, NAMEDB. You have put here your UserName, Password and DB name that you must create
before start work with this project. For quick start you may open MySql and write Query -> `create database usersdb` then another
one Query -> `create table Users ( id int auto_increment primary key, email varchar(30), activation_code varchar(30))`
  
#### 📌 How the project works

1. Press **Run** button or **^R** in Intellij IDEA to start server
2. Type in browser: http://localhost:8080/
3. In the **email** field and **retype email**, enter the email address. It should be noted
   that at this stage a simple check for empty fields has been implemented. If you leave some fields empty
   you will be informed about it. Also, if you try to register email that is already in DB - you will be informed 😎.
4. In next page you will see a value of Activation Code (this was made as example for simplicity of understanding).
When you will write again yours Email from previous step, you may choose write code that you see (RIGHT CODE) or 
as example make mistake and see what happens. Here realized check for empty fields, correct putting code in Activation
code field
5. When you put proposed Code you will see final page with information -
   Service started with email (your email) and has status: (status)



# golangEmailDemo
