package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func handleInput() {
	deltaX := rl.Vector3Subtract(camera3D.Target, camera3D.Position)
	deltaX.Y = 0
	deltaX = rl.Vector3Normalize(deltaX)
	deltaX = rl.Vector3Scale(deltaX, MOVE_SPEED)
	deltaZ := rl.NewVector3(deltaX.Z, 0, -deltaX.X)
	deltaX = rl.Vector3Scale(deltaX, rl.GetFrameTime())
	deltaZ = rl.Vector3Scale(deltaZ, rl.GetFrameTime())

	if rl.IsKeyDown(rl.KeyW) {
		camera3D.Position = rl.Vector3Add(camera3D.Position, deltaX)
		camera3D.Target = rl.Vector3Add(camera3D.Target, deltaX)
	}
	if rl.IsKeyDown(rl.KeyA) {
		camera3D.Position = rl.Vector3Add(camera3D.Position, deltaZ)
		camera3D.Target = rl.Vector3Add(camera3D.Target, deltaZ)
	}
	if rl.IsKeyDown(rl.KeyS) {
		camera3D.Position = rl.Vector3Subtract(camera3D.Position, deltaX)
		camera3D.Target = rl.Vector3Subtract(camera3D.Target, deltaX)
	}
	if rl.IsKeyDown(rl.KeyD) {
		camera3D.Position = rl.Vector3Subtract(camera3D.Position, deltaZ)
		camera3D.Target = rl.Vector3Subtract(camera3D.Target, deltaZ)
	}

	if rl.IsKeyDown(rl.KeySpace) {
		camera3D.Position.Y += ASCEND_SPEED * rl.GetFrameTime()
		camera3D.Target.Y += ASCEND_SPEED * rl.GetFrameTime()
	}
	if rl.IsKeyDown(rl.KeyLeftShift) {
		camera3D.Position.Y -= ASCEND_SPEED * rl.GetFrameTime()
		camera3D.Target.Y -= ASCEND_SPEED * rl.GetFrameTime()
	}

	if rl.IsKeyPressed(rl.KeyF11) {
		rl.ToggleFullscreen()
	}
	if rl.IsKeyPressed(rl.KeyF10) {
		if rl.IsCursorHidden() {
			rl.EnableCursor()
		} else {
			rl.DisableCursor()
		}
	}
	if rl.IsKeyPressed(rl.KeyF9) {
		SHOW_PLAYER_TARGET = !SHOW_PLAYER_TARGET
	}

	if rl.IsMouseButtonPressed(rl.MouseButtonLeft) {
		block, blockPos, chunk := world.getClosestBlockHit(getPlayerRay(), MAX_TRIGGERED_CHUNK_TARGET_SEARCH)
		if block != nil && chunk != nil && blockPos != nil {
			chunk.addBlock(AirBlock, worldPos3ToLocalPos3(*blockPos))
			chunk.generateChunkMesh()
		}
	}
}

func drawPlayerTarget() {
	block, blockPos, _ := world.getClosestBlockHit(getPlayerRay(), MAX_CONTINUOUS_CHUNK_TARGET_SEARCH)
	if block != nil && blockPos != nil {
		rl.DrawCube(rl.Vector3Add(position3ToVector3(*blockPos), rl.Vector3{X: 0.5, Y: 0.5, Z: 0.5}), 1.01, 1.01, 1.01, rl.Color{G: 121, B: 241, A: 127})
	}
}

func drawCrosshair() {
	rl.DrawTexture(cursor, int32(rl.GetScreenWidth()/2)-cursor.Width/2, int32(rl.GetScreenHeight()/2)-cursor.Height/2, rl.White)
}

func getPlayerRay() rl.Ray {
	return rl.GetScreenToWorldRay(rl.Vector2{X: float32(rl.GetScreenWidth()) / 2.0, Y: float32(rl.GetScreenHeight()) / 2.0}, camera3D)
}
