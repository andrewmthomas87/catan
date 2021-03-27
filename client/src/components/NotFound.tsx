import { Button } from '@chakra-ui/button'
import { Container, Heading, VStack } from '@chakra-ui/layout'
import React from 'react'
import { useLocation } from 'wouter'
import Nav from './Nav'

export default function NotFound() {
	const [_, setLocation] = useLocation()

	return (
		<>
			<Nav />
			<Container maxW="container.lg">
				<VStack spacing={2}>
					<Heading size="2xl">404</Heading>
					<Heading size="lg">Not found</Heading>
					<Button onClick={() => setLocation('/')}>Go home</Button>
				</VStack>
			</Container>
		</>
	)
}
