package main

import rl "github.com/gen2brain/raylib-go/raylib"

func (world *World) renderWorld(renderDistance int) {
	for chunkPos, chunk := range world.chunks {
		cameraWorldPos := Position2{X: int(camera3D.Position.X), Z: int(camera3D.Position.Z)}
		cameraChunkPos := worldPos2ToChunkPos2(cameraWorldPos)
		if chunkPos.X-cameraChunkPos.X <= renderDistance &&
			chunkPos.X-cameraChunkPos.X >= -renderDistance &&
			chunkPos.Z-cameraChunkPos.Z <= renderDistance &&
			chunkPos.Z-cameraChunkPos.Z >= -renderDistance {
			chunk.mesh.render()
		}
	}
	rl.DrawGrid(2*WORLD_SIZE, CHUNK_SIZE)
}

func (chunkMesh *ChunkMesh) render() {
	rl.DrawModel(chunkMesh.Model, rl.Vector3{}, 1.0, rl.White)
}
