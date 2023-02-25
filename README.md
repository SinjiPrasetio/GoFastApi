GoFastApi
=========

GoFastApi is a fast and easy-to-use framework for scaffolding Go lang APIs. This framework comes with pre-made helper functions that make it easy to get your API up and running in no time. Additionally, it provides a seamless integration with Docker to deploy your application to any environment.

Requirements
------------

Before getting started, ensure you have the following installed:

- Go version 1.16 or higher
- Make
- Docker

Getting Started
---------------

To use GoFastApi, simply clone this repository and start building your API. You can clone the repository by running:

bash

`git clone https://github.com/SinjiPrasetio/GoFastApi.git`

To build the application, simply run:

go

`make build`

This will build the application into a binary executable that you can find in the `build` directory. You can start your API by running:

bash

`./build/<APP_NAME>`

Replace `<APP_NAME>` with the name of the executable file, which is specified in the `.env` file.

Development
-----------

To run the application for development, you can use the following command:

go

`make run`

This will start the application in development mode, where the code changes will be automatically reloaded without the need to rebuild the application.

Docker Support
--------------

GoFastApi comes with built-in support for Docker to simplify the deployment process. To build the Docker image and start the containers, simply run:

go

`make start_compose`

This will build the Docker image for your application and start the containers. Your API will now be available on `http://localhost:8080`.

To stop the Docker image, run:

go

`make stop_compose`

This will stop the Docker containers and remove the Docker image.

License
-------

This code is released under the [MIT License](https://opensource.org/licenses/MIT). You are free to use, modify, and distribute the code as long as it complies with the terms and conditions of this license.

Conclusion
----------

GoFastApi is a simple and easy-to-use framework for scaffolding Go lang APIs. With its pre-made helper functions and seamless Docker integration, you can quickly and easily build and deploy your application. If you have any questions or feedback, please feel free to reach out. Happy coding!
