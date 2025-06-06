package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (world *World) colliderInit() {
	for chunkPos, chunk := range world.chunks {
		chunkMin := chunkAndLocalToWorldPos(Position{X: chunkPos.X, Z: chunkPos.Z}, Position{X: 0, Z: 0})
		chunkMax := chunkAndLocalToWorldPos(Position{X: chunkPos.X, Z: chunkPos.Z}, Position{X: CHUNK_SIZE, Z: CHUNK_SIZE})
		chunk.collider = rl.NewBoundingBox(rl.Vector3{X: float32(chunkMin.X), Y: 0, Z: float32(chunkMin.Z)}, rl.Vector3{X: float32(chunkMax.X), Y: CHUNK_HEIGHT, Z: float32(chunkMax.Z)})
		for x := range chunk.blocks {
			for z := range chunk.blocks[x] {
				for y, block := range chunk.blocks[x][z] {
					worldPos := chunkAndLocalToWorldPos(chunkPos, Position{X: x, Z: z})
					worldPosFloat32 := rl.Vector3{X: float32(worldPos.X), Y: float32(y), Z: float32(worldPos.Z)}
					block.collider = rl.NewBoundingBox(worldPosFloat32, rl.Vector3Add(worldPosFloat32, rl.Vector3{X: 1, Y: 1, Z: 1}))
				}
			}
		}
	}
}

func (world *World) getClosestBlockHit(ray rl.Ray, maxDistance float32) (*Block, *Chunk) {
	var closestBlock *Block = nil
	var closestChunk *Chunk = nil
	for chunkPos, chunk := range world.chunks {
		chunkCol := rl.GetRayCollisionBox(ray, chunk.collider)
		if chunkCol.Hit && rl.Vector3Distance(positionToVector3(chunkPos), positionToVector3(worldToChunkPos(vector3ToPosition(ray.Position)))) < maxDistance {
			for x := range chunk.blocks {
				for z := range chunk.blocks[x] {
					for _, block := range chunk.blocks[x][z] {
						blockCol := rl.GetRayCollisionBox(ray, block.collider)
						if closestBlock == nil || (blockCol.Hit && block.data != AirBlock &&
							rl.Vector3Distance(block.collider.Min, ray.Position) < rl.Vector3Distance(closestBlock.collider.Min, ray.Position)) {
							closestBlock = block
							closestChunk = chunk
						}
					}
				}
			}
		}
	}
	return closestBlock, closestChunk
}
