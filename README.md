# PrepTheSpire

A codebase that utilizes slay the spire tier scores for cards to help players slay the spire.
This codebase uses gqlgen to create a graphql server that display card tiers for slay the spire. With time, I'm hoping I can create a "run helper" app that uses the info from this graphql server to create a full fledged mobile app that everyone can use to make informed decisions about slay the spire.

# Setup

Clone the code base
Download Go for your OS
Create a mongodb database called SlayTheSpire
Create a collection called CardsTierList
Import the slay-the-spire.csv into the CardsTierList collection
Follow the setup below in resources

# Resources

I followed a tutorial here to get a simple setup:
https://www.section.io/engineering-education/implement-graphql-api-using-golang-and-mongodb-database/
