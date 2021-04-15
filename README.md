# E-SAFE

> A Distributed Storage Solution that allows for Identity and Access Management (IAM) and secure storage of Secrets while distributing security liability.​
> 50.041 Distributed Systems Project

## How to run?

To start as Locksmith Server:

```
go run cmd/e-safe/main.go -locksmith
```

To start as a Node:

```
go run cmd/e-safe/main.go -node
```

To start the front-end web application:

```
cd client
npm install
npm run serve
```

Visit the website at: `localhost:8080`

## Demo

![](./doc/demo.gif)
