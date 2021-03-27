import { Box } from '@chakra-ui/layout'
import { useToken } from '@chakra-ui/system'
import Panzoom from '@panzoom/panzoom'
import React, { useEffect, useRef } from 'react'
import Hex from './Hex'

export default function Game() {
	const boxRef = useRef<HTMLDivElement>(null)
	const [blue500] = useToken('colors', ['blue.500'])

	useEffect(() => {
		const svg = boxRef.current!.querySelector<SVGElement>('svg')
		const panzoom = Panzoom(svg!, {
			contain: 'outside',
			cursor: 'initial',
			panOnlyWhenZoomed: true,
			minScale: 1,
		})
		svg!.parentElement!.addEventListener('wheel', panzoom.zoomWithWheel)

		return () => {
			svg!.parentElement!.removeEventListener('wheel', panzoom.zoomWithWheel)
			panzoom.destroy()
		}
	}, [boxRef])

	const size = 20
	const width = (size * Math.sqrt(3)) / 2
	const height = size

	const points: number[][] = []
	for (let y = 0; y < 4; y++) {
		for (let x = 0; x < 4; x++) {
			points.push([width * (x + (y % 2 === 0 ? 0.5 : 0)), height * 0.75 * y])
		}
	}

	return (
		<Box ref={boxRef} width="full" height="full" overflow="hidden">
			<svg
				width="100%"
				height="100%"
				viewBox="0 0 100 100"
				xmlns="http://www.w3.org/2000/svg"
				style={{ backgroundColor: blue500 }}
			>
				{points.map((point, index) => (
					<Hex key={index} x={point[0] + width} y={point[1] + height} size={size} />
				))}
			</svg>
		</Box>
	)
}
