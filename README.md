# Payment API - Description

## Table of Contents

- [Table of Contents](#table-of-contents)
- [Description](#description)
  - [Microservices - Payment](#microservice---payment)
- [Architecture](#architecture)
  - [DDD](#ddd)
  - [Clean Archtecture](#clean-archtecture)
- [Design Patterns](#design-patterns)
- [Unit Testing](#unit-testing)
  - [BDD](#bdd)
- [Docker build and run](#docker-build-and-run)
- [How to use](#how-to-use)
  - [Check app status](#check-app-status)
- [AWS](#aws)
- [Kubernetes](#kubernetes)
- [Section 1 - Restaurant owner](#section-1-restaurant-owner)
  - [1 Product manipulation](#1-product-manipulation)
- [Section 2 Customer order](#section-2-customer-order)
  - [1 User identification](#1-user-identification)
  - [2 List all the categories](#2-list-all-the-categories)
  - [3 List products by the chosen category](#3-list-products-by-the-chosen-category)
  - [4 Pay the products amount](#4-pay-the-products-amount)
    - [4.1 Generate Mercado Livre QR Code](#4_1-generate-mercado-livre-qr-code)
  - [5 Create an order](#5-create-an-order)
  - [6 List orders to follow](#6-list-orders-to-follow)
  - [7 List orders to prepare](#7-list-orders-to-prepare)
  - [8 List orders waiting payment](#8-list-orders-waiting-payment)
  - [9 Update order to preparing](#9-update-order-to-preparing)
  - [10 Update order to done](#10-update-order-to-done)
  - [11 Update order to delivered](#11-update-order-to-delivered)
  - [12 Update order to not delivered](#12-update-order-to-not-delivered)
- [Mercado Livre Webhook](#mercado-livre-webhook)
- [Documentation](#documentation)
  - [Event storming](#event-storming)
  - [Postman collection](#postman-collection)
  - [Swagger](#swagger)
  - [Redoc](#redoc)

## Description

The Tech Challenge 1 aims to do a solution for a Fast Food restaurant. With this software, the rastaurant can do a all the process of for a place
that makes all steps of a fast food dish order, as:

- Products creation/manipulation by the restaurant owner
- Customer identification
- Order creation with given products
- Payment process
- - With Mercado Livre QR Code payment option [Webhook Payment](internal/core/webhook/README.md)
- Order tracking by the chef
- Order tracking by the waiter and the customer

This projects only fits the Backend side, which means that customer needs to **choose** the products or combo by a interface previously. This Backend will only receive the *entire order with all chosen products or combos*. This Backend will not do a *step by step product selecion*.

All the Endpoints can be called by accessing `http://localhost:3210/api` API url.

To build and run this project. Follow the Docker section

### Microservice - Payment

The Project is divided in `3 microservices`. Each one has its own Database, logic and POD inside `EKS Cluster`. The microservices are:

- Customer
- Orders
- Payment

This Microservice is responsible for the `Payment` API. This microservice uses `Altas Mongo DB` to save its data.

## Architecture

This project uses a mixin with two architectures to make it scalable and secure by protecting the **domain** layer. Those are `DDD (Domain Driven Design)` and `Clean Archtecture`

### DDD ###

To design this application was chosen `DDD (Domain Drive Design)` architecture to follow the principle of **protecting the model**.

![ddd_image](https://github.com/thiagoluis88git/tech1-payment/assets/166969350/2016bfff-3c19-4172-837f-8d5d428525f7)

### Clean Archtecture ###

The other one is `Clean Archtecture`. With it, we add some extra layers to organize even more the project.

![CleanArchitecture](https://github.com/user-attachments/assets/a49c2aab-562c-4b6c-82f2-7ffe9e4aec74)

The folder project was created to follow this main principle:

- **data**: Here we have all the implementations, such as Repositories, Remotes and Locals
- **domain**: All the `application business logic`, such as UseCases
- **handler**: Also known as `presentation layer` resides all the `Controllers` handled by the `Web` **interface** given by the `/cmd/api Framework & Driver` 

- **integrations**: This is not part of `DDD` or `Clean Arch`, but is important separate some external packages or integrations, such `Mercado Livre API`

## Design Patterns

To improve and make a good standard project pattern, some `Design Patterns` were used in this application.

- Strategy: All the business logic must be protected by the external implementations. To do it, we use a combo with **Interfaces** and **Dependency Inversion solid principle** to inject only *interfaces* and not *real implementations*
- Dependency Injection: Is used in application bootstrap (main.go) to inject all the interfaces implementations.
- Decorator: To inject **Services** inside **Driver Adapter** *handler*. By doing this, we *decorate* the Handler with a Service
- Services or Use Cases: Centralize all the business logic of of the application
- Repository: Used to integrate with all **Driven Adapter** like *Databases and External Endpoints*

## Unit Testing

To run all the Unit Testing for this project, just run:

```
go test -cover ./... -coverprofile="cover.out"
go tool cover -func="cover.out"
```

This will run all the **Services** unit tests and **Repository** unit Database tests running [Testcontainers](https://testcontainers.com/) database container mocks.

### BDD

Inside `bdd` folder, has an implementation of a BDD test. This test is made by a [Cucumber API](https://github.com/cucumber/godog).
The **BDD** tests will be triggered when running the `go test ./...` in the previous step

```
go test -cover ./... -coverprofile="cover.out"
go tool cover -func="cover.out"
```

## Docker build and run

This project was built using Docker and Docker Compose. So, to build and run the API, we need to run in the root of the project:

```
docker compose build
```

After the image build finish, run:

```
docker compose up -d
```

The command above may take a while.

After all the containers shows these below status:

```
 ✔ Container fastfood-database  Started
 ✔ Container fastfood-app       Started 
```

we can access `http://localhost:5210/api` endpoints.


## How to use

To use all the endpoints in this API, we can follow these sequence to simulate a customer making an order in a restaurant.
We can separate in three moments.

- Restaurant products manipulation. This is used by the `restaurant owner` to create all the product portfolio with its images and prices
- Customer self service. This is used by the `customer` to choose the products, pay for it and create an order 
- Order preparing and deliver. This is used by the `chef` and `waiter` to check the order status

We will divide in 2 sections: **Restaurant owner** and **Customer order**

### Check app status

After running `Docker` commands, you can check the application status running:

```
docker compose logs app
```

We can see some database errors but at the end of the logs we can see:

```
fastfood-app  | 2024/05/27 22:57:35 API Tech 1 has started
```

## AWS ##

The Fast food project uses `AWS Cloud` to host its software components. To know more about the **AWS configuration**, read: [AWS Readme](https://github.com/thiagoluis88git/tech1-k8s/infra/README.md)

## Kubernetes

This application has all the K8S YAMLs to be applied in any cluster. 
To read the specific documentation, read: [Kubernetes README](https://github.com/thiagoluis88git/tech1-k8s/infra/k8s/README.md)

## Section 1 Restaurant owner

This section will be used by the restaurant owner to manage the restaurant products

### 1 Product manipulation
***(Owner view)***

See [Orders README](https://github.com/thiagoluis88git/tech1-orders/blob/main/README.md)


## Section 2 Customer order

This section will use all the Endpoints to make a entire order flow.

### 1 User identification
***(Customer view)***

See [Customer README]()

### 2 List all the categories
***(Customer view)***

See [Orders README](https://github.com/thiagoluis88git/tech1-orders/blob/main/README.md)

### 3 List products by the chosen category
***(Customer view)***

See [Orders README](https://github.com/thiagoluis88git/tech1-orders/blob/main/README.md)

### 4 Pay the products amount
***(Customer view)***

- Call the GET `http://localhost:5210/api/payments/types` to show to customer which payment type to choose
- Call the POST `http://localhost:5210/api/payments` to pay for the amount and receive the `[Payment ID]`

#### 4_1 Generate Mercado Livre QR Code ####
***(Customer view)***

- Call the POST `http://localhost:5210/api/qrcode/generate` to get the `QR Code Data` to **transform** in Image to pay with `Mercado Pago App`. Must send the same **post body** as [5 create an order](#5-create-an-order) needs.

> [!WARNING]
> Sometimes the **Mercado Livre** server returns `500 Internal Server Error` for unknown reason. The error returned by the server is: `{"error":"alias_obtainment_error","message":"Get aliases for user failed","status":500,"causes":[]}`. When this occurs, **IS NOT possible to proceed with QR Code Payment**. The main reason for this is on `Weekend the Mercado Livre development environment does not work`

When the server returns as expected, the response is like:

```
{
    "data": "00020101021243650016COM.MERCADOLIBRE020130636b624f3f3-d289-453b-a620-a32198e16a235204000053039865802BR5909Test Test6009SAO PAULO62070503***63046773"
}
```

> [!NOTE]
> This method will `CREATE AN ORDER` with the status `WAITING PAYMENT`.

### 5 Create an order
***(Customer view)***

See [Orders README](https://github.com/thiagoluis88git/tech1-orders/blob/main/README.md)

### 6 List orders to follow
***(Customer and Waiter)***

See [Orders README](https://github.com/thiagoluis88git/tech1-orders/blob/main/README.md)

### 7 List orders to prepare
***(Chef view)***

See [Orders README](https://github.com/thiagoluis88git/tech1-orders/blob/main/README.md)

### 8 List orders waiting payment
***(Owner view)***

See [Orders README](https://github.com/thiagoluis88git/tech1-orders/blob/main/README.md)

### 9 Update order to preparing
***(Chef view)***

See [Orders README](https://github.com/thiagoluis88git/tech1-orders/blob/main/README.md)

### 10 Update order to done
***(Chef view)***

See [Orders README](https://github.com/thiagoluis88git/tech1-orders/blob/main/README.md)

### 11 Update order to delivered
***(Waiter view)***

See [Orders README](https://github.com/thiagoluis88git/tech1-orders/blob/main/README.md)

### 12 Update order to not delivered
***(Waiter view)***


## Mercado Livre Webhook ##

The Fast Food application can pay the order via QR Code. 
This is a separate flow and can be read in: [Webhook Payment](internal/core/webhook/README.md)

## Documentation

This project uses Swagger to show an site with all Endpoints used by this project to make an order in a Fast Food place. 
To create/update all Endpoints documentation just run `swag init -g cmd/api/main.go`. By doing this, we can see the documentation in
two different ways:

### Event Storming

This project was guided by the DDD Event Storming. This document was made in Miro. The file `event_storming.pdf` is in the root of this project.
We can algo see the Miro project by accessing the link:[Event Storming](https://miro.com/app/board/uXjVKL0pb-w=/)

### Postman collection

In the root of this project we can find the file `postman_collection.json`. With this we can easly test all the Endpoints. It also be in this repo all the `postman environments` to use in `localhost` and  with `EKS LoadBalancer DNS`. Just import also these:
- eks_postman_environment.json
- localhost_postman_environment.json
- minikube_postman_environment.json

### Swagger

http://localhost:5210/swagger/index.html

### Redoc

http://localhost:5211/docs
