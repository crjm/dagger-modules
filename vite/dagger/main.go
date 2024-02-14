package main

import (
	"context"
)

type Vite struct{}

// https://v3.vitejs.dev/config/server-options.html#server-port
const defaultPort int = 5173
const defaultCacheLocation = "/root/.npm"

// prints the contents of the vite project directory
func (m *Vite) Source(dirPath *Directory) *Container {
  dir := dag.Container().From("node:lts").WithMountedDirectory("app", dirPath).WithWorkdir("app").WithExec([]string{"ls"})
  return dir
}

// install project dependencies with npm
func (m *Vite) InstallDeps(ctx context.Context, dirPath *Directory) *Container {
  cache := dag.CacheVolume("npm-cache")
  return m.Source(dirPath).WithMountedCache(defaultCacheLocation, cache).WithExec([]string{"npm", "install"})
}

// builds the project and saves the output to /app/dist dir in the container
func (m *Vite) Build(ctx context.Context, dirPath *Directory) *Container {
  return m.InstallDeps(ctx, dirPath).WithExec([]string{"npm", "run", "build"})
}

