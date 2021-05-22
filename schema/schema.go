package schema

const Schema = `
	schema {
		query: Query
	}
	
	type Query {
		// user(id: ID!): User!
		product(id: ID!): Product!
		products(): [Product]!
		carts(): [Cart]!
		addToCart(productID: Int!, userID: Int!, quantity: Int!): Cart!
		checkout(userID: Int!, pay: Float!): Checkout!
	}

	type Checkout {
		totalOrder: Float!
		message:    String!
		isSuccess:  Boolean!
	}

	type Product {
		sku: String!
		name: String!
		price: Float!
		promoCode: String!
		quantity: Int!
	}

	type Cart {
		userId: Int!
		productId: Int!
		total: Float!
		quantity: Int!
		description: String!
		freeItems: Int!
		totalFreeItems: Int!
		status: String!
	}

`

type Resolver struct{}
