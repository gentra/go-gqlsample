package internal

const (
	Schema = `
				schema {
					query: Query
					mutation: Mutation
				}

                type Query {
					hello: String!

					# A list of all Lift objects
					allLifts(status: LiftStatus): [Lift!]!
                }

				type Mutation {
					setLiftStatus(id: ID!, status: LiftStatus!): Lift!
				}

				# A Lift is a chairlift, gondola, tram, funicular, pulley, rope tow, or other means of ascending a
				# mountain.
				type Lift {
					id: ID!
					name: String!
					status: LiftStatus
					capacity: Int!
					night: Boolean!
					elevationGain: Int!

					# A list of trails that this Lift serves
					trailAccess: [Trail!]!
				}

				# A Trail is a run at a ski resort
				type Trail {
					id: ID!
					name: String!
					status: TrailStatus
					difficulty: String!
					groomed: Boolean!
					trees: Boolean!
					night: Boolean!
					#accessedByLifts: [Lift!]!
				}

				# An enum describing the options for LiftStatus: OPEN, CLOSED, HOLD
				enum LiftStatus {
					OPEN
					CLOSED
					HOLD
				}

				# An enum describing the options for TrailStatus: OPEN, CLOSED
				enum TrailStatus {
					OPEN
					CLOSED
				}
`
)
