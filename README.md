# 

how to start the services:
- open your terminal and do preferably inside /home/<username>/go/* path: git clone https://github.com/krabhanand/sellersappassignment.git
- next cd into sellersappassignment: cd sellersappassignment
- make sure docker and docker-compose are installed
- do: sudo docker-compose up --build
- wait for everthing to startup and some json data displayed
- have a look at mongodb data first, open new terminal and look up existing docker containers, type: docker ps
- copy the container id whose image name is 'mongo:latest'
- type: docker exec -it <container-id> bash
- now we are inside docker container that has mongodb server running
- type : mongo
- now we are in mongodb shell
- type: use seller_app
- type: db.product_data.insertOne({"a":"q"})
- now data base has been initialized
  
now test the service
- go to postman, type in address "http://localhost:10000/"
- in params tab enter 'url' as a parameter key and """ https://www.amazon.com/PlayStation-4-Pro-1TB-Console/dp/B01LOP8EZC/  """ as its value(type only the web address, do not type triple inverted commas)(you may type any other amazon.com products web page url)
- make a post request
- api functions as given below and relevant data is returned after being stored in the database
- go to the mongo shell and type : db.product_data.find()
- you can see the data extracted from the web page of enetered url of amazon products page
  
  
  
two rest api have been created as mentioned in the assignment document
db used here is mongodb with it being packaged along with docker
the name of the database is 'seller_app', the name of collection is 'product_data'
first api is present in folder appseller,
  - the API in here runs on port 10000.
  - it receives a URL in the query parameter 'url'
  - if the query paramater 'url' is not present, relevant error message will be sent along with error code 400
  - if the url would not belong to amazon product page or would not be a valid url relevant error message would be sent along with 400 code
  - after receving the url, the api would do scraping of data from web page using goQuery
  - scraping code is written in a different file for better maintainance of code
  - as a result of scraping: product is available containg product Name, Description, Price, Reviews, ImageURL - given 5 feilds
  - this product feild along with url is sent to the second api as a post request
  - after receiving response as ok, the data saved in the db(barring the id) along with timestamp which is obtained from the other API is provided to the user for satisfaction

second api is present in the 'db-save-scrap-data' folder,
  - the API runs on port 8080
  - it receives JSON data from the first API regarding url and relevant details of the product
  - fresh timestamp is inserted in the model data
  - at first the db is searched for if data with same url has been stored before
  - if yes, then document is updated, if no, then new document is inserted in the collection
  - if every thing goes on fine, then json of inserted data(with dummy id) is provided back to the first API
  - if errors occour, then 500 error code along with relevant details is provided
  
both apis have been dockerized and canbe started with docker-compose up command
