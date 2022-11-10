# PrepTheSpire

![study](hero_studying.jpg)

A codebase that utilizes slay the spire tier scores for cards to help players slay the spire.
This codebase uses gqlgen to create a graphql server that display card tiers for slay the spire. With time, I'm hoping I can create a "run helper" app that uses the info from this graphql server to create a full fledged mobile app that everyone can use to make informed decisions about slay the spire.

# To do list

- Debug why the act1, act2, act3, overallScore, and upgradeBonus return 0...

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
