import { useAuth0 } from '@auth0/auth0-react'
import { Alert, AlertDescription, AlertIcon, AlertTitle } from '@chakra-ui/alert'
import { Container } from '@chakra-ui/layout'
import { useObservable, useObservableState } from 'observable-hooks'
import React from 'react'
import { interval } from 'rxjs'
import { map, startWith, take, tap } from 'rxjs/operators'

const REDIRECT_WAIT_TIME = 3

export default function SignedOut() {
	const { loginWithRedirect } = useAuth0()

	const countdown$ = useObservable(() =>
		interval(1000).pipe(
			map((value) => REDIRECT_WAIT_TIME - value - 1),
			take(REDIRECT_WAIT_TIME),
			startWith(REDIRECT_WAIT_TIME),
			tap((value) => {
				if (value === 0) {
					loginWithRedirect()
				}
			})
		)
	)
	const countdown = useObservableState(countdown$, REDIRECT_WAIT_TIME)

	return (
		<Container maxW="container.lg">
			<Alert status="info">
				<AlertIcon />
				<AlertTitle>Signed out</AlertTitle>
				<AlertDescription>Redirecting to login in {countdown}...</AlertDescription>
			</Alert>
		</Container>
	)
}
