import { Auth0Provider } from '@auth0/auth0-react'
import { ChakraProvider, ColorModeScript, extendTheme } from '@chakra-ui/react'
import React from 'react'
import ReactDOM from 'react-dom'
import Root from './components/Root'
import { config } from './config'
import './css/global.css'

const theme = extendTheme({
	config: { initialColorMode: 'dark' },
})

ReactDOM.render(
	<React.StrictMode>
		<ColorModeScript initialColorMode={theme.config.initialColorMode} />
		<ChakraProvider theme={theme}>
			<Auth0Provider
				domain={config.auth0.domain}
				clientId={config.auth0.clientID}
				audience={config.auth0.audience}
				redirectUri={config.auth0.redirectURI}
			>
				<Root />
			</Auth0Provider>
		</ChakraProvider>
	</React.StrictMode>,
	document.getElementById('root')
)
