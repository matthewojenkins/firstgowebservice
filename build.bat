
rem Go Build
set GOOS=linux
set GOARCH=amd64

cd src
go build -o ..\build\program .

rem Docker Build
cd ..
docker build -f Dockerfile -t programservice:1.0 .


rem Upload to GCP container repository
rem Get image id
rem docker image ls

rem docker tag 63d164e09f3e eu.gcr.io/workoutapp-270814/programservice
rem gcloud auth login
rem gcloud auth configure-docker
rem docker push eu.gcr.io/workoutapp-270814/programservice


rem Run docker locally
docker run -p 8080:8080 --name programservice programservice:1.0