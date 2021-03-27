import { Button } from '@chakra-ui/button'
import Icon from '@chakra-ui/icon'
import { AddIcon } from '@chakra-ui/icons'
import { Container } from '@chakra-ui/layout'
import React from 'react'
import { useLocation } from 'wouter'
import Nav from './Nav'

export default function Home() {
	const [_, setLocation] = useLocation()

	return (
		<>
			<Nav />
			<Container maxW="container.lg">
				<Button rightIcon={<Icon as={AddIcon} />} onClick={() => setLocation('/game')}>
					New game
				</Button>
			</Container>
		</>
	)
}
