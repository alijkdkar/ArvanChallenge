# ArvanChallenge



##this practice include two micro services 
#1- core Service for manage user and cards 
  - core service use of postgres as Database
  - Orm is gorm
  - reids package for connect with Redis message bus


#2- micro service for mange discount opportunity (folder start with "discount")
  - discount service use of redis as Datebase
  - use of redis ORM for connect to Redis-CLI
  - use of redis package for connect redis cli message bus
    

#tow services have relatin with redis message bus 
this practice use of "out box pattern" for integrated transaction with core service


both of this service use of gin packeage for endpoint provider.

the rest api end point list is :

GET    /user/all                 --> /ArvanChallenge/controller.getAll (3 handlers)
POST   /user                     --> /ArvanChallenge/controller.createUser (3 handlers)
DELETE /user/:id                 --> /ArvanChallenge/controller.deleteUser (3 handlers)
GET    /user/:id                 --> /ArvanChallenge/controller.getUserDetail (3 handlers)
PUT    /user/:id                 --> /ArvanChallenge/controller.updateUser (3 handlers)
GET    /user/discountused/:Id    --> /ArvanChallenge/controller.GetUsersUsedDiscount (3 handlers)
GET    /card/all/:Id             --> /ArvanChallenge/controller.getAllCards (3 handlers)
POST   /card/:Id                 --> /ArvanChallenge/controller.createCard (3 handlers)
GET    /card/:userId/detail/:Id  --> /ArvanChallenge/controller.getCardDetail (3 handlers)
DELETE /card/:Id                 --> /ArvanChallenge/controller.deleteCard (3 handlers)
PUT    /card/:userId/update/:Id  --> /ArvanChallenge/controller.updateCard (3 handlers)
POST   /card/:Id/Transaction     --> /ArvanChallenge/controller.CreateTransaction (3 handlers)
GET    /card/Transaction/:cardId --> /ArvanChallenge/controller.getCardTransactions (3 handlers)
POST   /discount                 --> /ArvanChallenge/discountController.CreateDiscountOppurtunity (3 handlers)
DELETE /discount/:Id             --> /ArvanChallenge/discountController.DeleteDiscount (3 handlers)
[GIN-debug] POST   /discount/:Id/use/:cardId --> github.com/alijkdkar/ArvanChallenge/discountController.UseOfDiscount (3 handlers)
