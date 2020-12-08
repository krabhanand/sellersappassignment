# sellersappassignment
two rest api have been created as mentioned in the assignment document
db used here is mongodb with it being on the cloud and can be accessed via compass by the use of the link: mongodb+srv://anand:anand@cluster0.rlk7s.mongodb.net/
the name of the database is 'seller-app', the name of collection is 'product-data'
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
  
 YET TO DO:
 - dockerize both the services
 - combining them both using docker compose file
 
 ISSUES BEING FACED:
  - to run docker locally, it seems the requirement is of windows 10 pro edition which is not currently available
  - due to above becoming difficult to complete the docker portion of the assignment
  - wish to clarify wether the 2 services need to be converted into microservices or should be lefty as is?
