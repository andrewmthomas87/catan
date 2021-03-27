import { config } from '@/config'
import { useAuth0 } from '@auth0/auth0-react'
import { Avatar } from '@chakra-ui/avatar'
import { IconButton } from '@chakra-ui/button'
import { AtSignIcon, LockIcon } from '@chakra-ui/icons'
import { Container, Flex, Heading, Link } from '@chakra-ui/layout'
import { Menu, MenuButton, MenuDivider, MenuItem, MenuList } from '@chakra-ui/menu'
import React from 'react'
import { Link as WouterLink, useLocation } from 'wouter'

export default function Nav() {
	const { user, logout } = useAuth0()
	const [_, setLocation] = useLocation()

	const handleLogout = () => logout({ returnTo: config.auth0.redirectURI })

	return (
		<Container maxW="container.lg">
			<Flex justify="space-between" align="center" height={20}>
				<Heading size="lg">
					<Link as={WouterLink} href="/">
						catan
					</Link>
				</Heading>
				<Menu placement="bottom-end">
					<MenuButton as={IconButton} icon={<Avatar name={user.name} size="sm" />} isRound />
					<MenuList>
						<MenuItem icon={<AtSignIcon />} onClick={() => setLocation('/profile')}>
							My profile
						</MenuItem>
						<MenuDivider />
						<MenuItem icon={<LockIcon />} onClick={handleLogout}>
							Logout
						</MenuItem>
					</MenuList>
				</Menu>
			</Flex>
		</Container>
	)
}
