import Hexagon from '@/utility/Hexagon'
import React, { useState } from 'react'

type Props = {
	x: number
	y: number
	size: number
}

export default function Hex({ x, y, size }: Props) {
	const [isHovered, setIsHovered] = useState(false)

	const points: number[][] = []
	for (let i = 0; i < 6; i++) {
		points.push([x + size * Hexagon.POINT_X[i], y + size * Hexagon.POINT_Y[i]])
	}

	return (
		<polygon
			points={points.map((p) => p.join(',')).join(' ')}
			fill={isHovered ? 'red' : 'white'}
			stroke="black"
			strokeWidth={size * 0.01}
			onMouseOver={() => setIsHovered(true)}
			onMouseOut={() => setIsHovered(false)}
		/>
	)
}
