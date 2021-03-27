export const config = {
	auth0: {
		domain: import.meta.env.SNOWPACK_PUBLIC_AUTH0_DOMAIN,
		clientID: import.meta.env.SNOWPACK_PUBLIC_AUTH0_CLIENT_ID,
		audience: import.meta.env.SNOWPACK_PUBLIC_AUTH0_AUDIENCE,
		redirectURI: import.meta.env.SNOWPACK_PUBLIC_AUTH0_REDIRECT_URI,
	},
}
