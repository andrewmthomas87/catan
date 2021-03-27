import React from 'react'
import { Route, Switch } from 'wouter'
import Game from './game/Game'
import Home from './Home'
import NotFound from './NotFound'

export default function SignedIn() {
	return (
		<Switch>
			<Route path="/" component={Home} />
			<Route path="/game" component={Game} />
			<Route component={NotFound} />
		</Switch>
	)
}
