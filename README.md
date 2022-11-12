# PrepTheSpire

![study](hero_studying.jpg)

A codebase that utilizes slay the spire tier scores for cards to help players slay the spire.
This codebase uses gqlgen to create a graphql server that display card tiers for slay the spire. With time, I'm hoping I can create a "run helper" app that uses the info from this graphql server to create a full fledged mobile app that everyone can use to make informed decisions about slay the spire.

# Notes

- 11-10-22
  I think my server.go code is not correct. The imports don't seem right. -update- Just kidding, it wasn't the server.go code. It was the
- 11-11-22
  Figured out what was wrong. I just needed to convert the object to a primitive map.

# To do list

- Add rarity, deck goals to csv

# Development flow for a new graphql route

- Define the schema.graphql file
- Run `go run github.com/99designs/gqlgen generate`
- Find and implement the method in database.go
- Update schema.resolvers.go to use said method from database.go

# Setup

- Clone the code base
- Download Go for your OS
- Create a mongodb database called SlayTheSpire
- Create a collection called CardsTierList
- Import the slay-the-spire.csv into the CardsTierList collection
- Follow the setup below in resources

# Starting the server

`go run server.go`

# Resources

I followed a tutorial here to get a simple setup:

- https://www.section.io/engineering-education/implement-graphql-api-using-golang-and-mongodb-database/
- https://reactnative.dev/docs/environment-setup
- https://www.stackhawk.com/blog/golang-cors-guide-what-it-is-and-how-to-enable-it/
- https://pkg.go.dev/go.mongodb.org/mongo-driver@v1.10.3/mongo#Collection.Find
- https://www.mongodb.com/docs/drivers/go/current/fundamentals/bson/
