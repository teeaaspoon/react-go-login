# react-go-login

This application is a login form using Gin, React and Postgres.
Sessions are managed with JWT and CSRF Tokens to prevent security vulnerabilities.
This allows only authenticated users to see certain pages and request data from the backend.

!["login_page"](https://github.com/teeaaspoon/react-go-login/blob/main/images/login_page.png)
!["authentication_flow"](https://github.com/teeaaspoon/react-go-login/blob/main/images/authentication_flow.png)

## Onboarding and Setup
Using docker-compose we can run all 3 services with one command

Prerequisites:
- [docker-compose](https://docs.docker.com/compose/install/)

```bash
# clone this repo
docker-compose build
docker-compose up
```
Your frontend should now be running on https://localhost and the backend will be running on 
https://localhost:8000

When you navigate to https://localhost you may see a security warning from your browser. This is because we
are using a self-signed certificate taken from the [fakeiot repo](https://github.com/gravitational/fakeiot/tree/master/fixtures).
Click Advanced and proceed to localhost.

!["untrusted_certificate"](https://github.com/teeaaspoon/react-go-login/blob/main/images/untrusted_certificate.png)


