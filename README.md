# <a href="https://tourdeapp.cz/" target="_blank">Tour de App</a>
- <a href="https://tourdeapp.cz/vzdelavaci-materialy" target="_blank">Materiály</a>
- <a href="https://tourdeapp.cz/zadani/introduction" target="_blank">Zadání</a>
- <a href="https://tinyurl.com/54ekrruk" target="_blank">Tour de Cloud</a>

## How to run

### Local development
- server:
    - in /apps/server/cmd/tourbackend run: go run .
    - runs on port 3000
    - if u change something you have to rerun the command
- web
    - in /apps/web run: npm run dev
    - app is automatically reloaded when you change something

### Local deployment with docker
This is the apps as they will be run in the cloud
After each change you'll have to rebuild the images and rerun the containers

- server:
    - in /apps/server run: docker build -t tourbackend .
    - this will create a docker image
    - to run this docker image ei create a container:
        - docker run --rm -p 3000:3000 --name tourbackend tourbackend
- web:
    - in /apps/web run: docker build -t tourfrontend .
    - this will create a docker image
    - to run this docker image ei create a container:
        - docker run --rm -p 3001:3001 --name tourfrontend tourfrontend

## Deploy to Tour de Cloud
git push automatically deploys to the cloud with a github action - once the fee is paid
as the fee hasn't been payed yet this is untested