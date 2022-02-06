# GoPersonalProjectsWIP
GoPersonalProjectsWIP


Some personal projects I've been working on .... still not done. Time permitting :-)


## ImageBox:

Is an an Image Upload service that runs in CloudRun. It will upload an image to GCP cloud-bucket and attach meta-data to the Cloud Firestore Metadata. 

TODO: Finish front-end, not a front-end guy so borrowing front-end design from different books. Use Hugo perhaps?  
TODO: move static files to either GCP Bucket or look at embedding them via the Go `embed` package that is available from Go 1.16. 
TODO: CI/CDI - for now updates CloudRun via manual deployment using gcloud cli (works). 
TODO: CLI for CLI upload??? 
TODO: Restructure server start. 
TODO: Better signalling shutdown and load shedding. 


## Telegrambot

Telegram Bot that will interact with CoinGecko API. For now project abandoned. 

TODO: lots :-)


## CloudCLI

See Readme in dir. Abandoned for now ....

## Change log

- Created repo
- Moved projects from various folders to this new repo.
