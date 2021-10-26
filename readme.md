# Finder
## _A simple powerful tool for finding things_

Finder is an mp3 search engine, built on Go with a ReactJS frontend.

This is an idea I've had for a long time and decided to throw together a small POC.

Right now the tool is limited to mp3 files and it works surprisingly well. I have plans to continue development and include many more features, this might eventually include other file formats but for now mp3 will do.

## It's ugly but it works

The front and backend are now conected and everything is working as a rough prototype.

Next up is to tidy how the results are displayed. Then I will finish the scraper tool to return more exact results.

## How it works

Right now you can search for an artist or specify an artist with song title. Finder will retrun a list of URLs which should contain your desired mp3.

Functionality and UI will be updated over the few commits.


## Setup

Run 'yarn install'
Then in seperate terminal windows, run these commands
'yarn start'
'go run .'

You should reach finder on 127.0.0.1:3000
