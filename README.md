# BlogApplication

### Purpose
The purpose of the app is to aggregate many blogs that provide RSS feed.
Then the app run the scraper concurrently with serving the REST api to serve the up-to-date blog posts for users.

### Features
- Authenticate users to create certain blogs that they want, and follow others' blogs.
- Scrape the metadata of blogs using Go concurrency.
- Using PostgresSQL as database.
