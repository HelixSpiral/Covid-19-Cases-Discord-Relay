# README
![BuildAndDeploy](https://github.com/HelixSpiral/Covid-19-Cases-Discord-Relay/workflows/BuildAndDeploy/badge.svg)
---

This is a simple script that relays the current Covid-19 case count to a Discord webhook.

Current case count is taken from [Here](https://coronavirus-19-api.herokuapp.com/all)

This is build for use with AWS Lambda.

To run outside of Lambda swap the comments in the main function, remove the Lambda import, and set the DISCORD_WEBHOOK environment variable to your webhook url.

Example usage outside of Lambda: `DISCORD_WEBHOOK="http://test.url/hook" go run main.go`