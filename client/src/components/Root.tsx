import { useAuth0 } from '@auth0/auth0-react'
import { Alert, AlertDescription, AlertIcon, AlertTitle } from '@chakra-ui/alert'
import { Center, Container } from '@chakra-ui/layout'
import { Spinner } from '@chakra-ui/spinner'
import React from 'react'
import SignedIn from './SignedIn'
import SignedOut from './SignedOut'

export default function Root() {
	const { isLoading, error, isAuthenticated } = useAuth0()

	if (isLoading) {
		return <Loading />
	} else if (error) {
		return <Error error={error} />
	}

	return isAuthenticated ? <SignedIn /> : <SignedOut />
}

function Loading() {
	return (
		<Center w="100%" h="100%">
			<Spinner />
		</Center>
	)
}

function Error({ error }: { error: Error }) {
	return (
		<Container maxW="container.lg">
			<Alert status="error">
				<AlertIcon />
				<AlertTitle>{error.name}</AlertTitle>
				<AlertDescription>{error.message}</AlertDescription>
			</Alert>
		</Container>
	)
}
