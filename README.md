# Rafael Store
To run the application after installing the latest version of go available at https://golang.org/dl/

**Stacks**: *Postgres, GoLang, Docker, SQL, HTML*

* To run the application it is necessary to run the simple **command**: 
  ```sh
  docker-compose up -d --build
  ```

The docker will download the dependencies and start the applications in order:
1. Pull and start postgres latest;
2. Run SQL command for create a database, table and insert initial rows;
3. Pull and start golang latest;
4. Run app if database is up;

**Application:**

To acess the application:
  ```sh
  http://localhost:8009/
  ```
The application is a simple product crud where you can insert edit and delete products 

<!-- CONTRIBUTING -->
## Contributing

Contributions are what make the open source community such an amazing place to be learn, inspire, and create. Any contributions you make are **greatly appreciated**.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some Amazing Feature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- CONTACT -->
## Contact

Rafael Loureiro Wagner - [@rafael.wagner9](https://www.linkedin.com/in/rafael-loureiro-wagner/) - rafaelwagner09@hotmail.com

Project Link: [https://github.com/rafaelwagner/go-rafaelstore](https://github.com/rafaelwagner/go-rafaelstore)
