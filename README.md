# URL Shortener

This is a simple URL shortener project built using the Go programming language and the Gin web framework. The application allows users to shorten long URLs into shorter, more manageable links, making it easier to share and remember.

## Description

URL shorteners are handy tools that take long URLs and convert them into shorter, unique links that redirect to the original URL when accessed. This project uses a combination of PostgreSQL for database storage and Redis for caching to efficiently manage the URLs.

The application provides a user-friendly web interface where users can enter their long URLs and receive shortened links that redirect to the original URLs when clicked. Additionally, the application performs caching of recently shortened URLs to optimize redirection speed.

## Features

- Shorten long URLs into unique and easily shareable links.
- Redirect users to the original URL when accessing the shortened link.
- Efficiently cache recently shortened URLs using Redis for faster redirection.
- Simple and intuitive web interface.
- Utilizes Google's UUID package to generate unique short URLs.

## Usage

Once the URL shortener is up and running, you can access it through your web browser or use the provided API endpoints programmatically.

**1. Access the Web Interface:**

Open your web browser and go to `localhost:8080` to access the URL shortener's web interface.

**2. Shorten a URL:**

Enter a long URL that you want to shorten, and the application will generate a unique short link.

**3. Redirect to Original URL:**

Click on the shortened link to be redirected to the original URL.


## Usage

Once the URL shortener is up and running, you can access it through your web browser or use the provided API endpoints programmatically.

1. **Access the Web Interface:**
   
   Open your web browser and go to `localhost:8080` to access the URL shortener's web interface.

2. **Shorten a URL:**
   
   Enter a long URL that you want to shorten, and the application will generate a unique short link.

3. **Redirect to Original URL:**
   
   Click on the shortened link to be redirected to the original URL.

