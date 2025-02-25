package tests

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/testcontainers/testcontainers-go"
)

var Debian = struct {
	DOCKER_IMAGE       string
	DOCKER_IMAGE_TAG   string
	AWS_ECR_URI        string
	DOCKER_IMAGE_GROUP string
}{
	DOCKER_IMAGE:       "debian",
	DOCKER_IMAGE_TAG:   "latest",
	AWS_ECR_URI:        "public.ecr.aws/w2u0w5i6",
	DOCKER_IMAGE_GROUP: "base",
}

func TestContainerBuildDebian(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{
				Context:       "../" + Debian.DOCKER_IMAGE + "/",
				Dockerfile:    "Dockerfile.debian",
				KeepImage:     false,
				PrintBuildLog: true,
			},
		},
		Started: true,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}

func TestContainerBuildDebianSlim(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			FromDockerfile: testcontainers.FromDockerfile{
				Context:       "../" + Debian.DOCKER_IMAGE + "/",
				Dockerfile:    "Dockerfile.debian-slim",
				KeepImage:     false,
				PrintBuildLog: true,
			},
		},
		Started: true,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}

func TestContainerPullDebian(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: Debian.AWS_ECR_URI + "/" + Debian.DOCKER_IMAGE_GROUP + "/" + Debian.DOCKER_IMAGE + ":" + Debian.DOCKER_IMAGE_TAG,
		},
		Started: false,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}

func TestContainerPullDebianSlim(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: Debian.AWS_ECR_URI + "/" + Debian.DOCKER_IMAGE_GROUP + "/" + Debian.DOCKER_IMAGE + ":" + Debian.DOCKER_IMAGE_TAG,
		},
		Started: false,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}

func TestContainerExecDebian(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: Debian.AWS_ECR_URI + "/" + Debian.DOCKER_IMAGE_GROUP + "/" + Debian.DOCKER_IMAGE + ":" + Debian.DOCKER_IMAGE_TAG,
			Cmd:   []string{"echo", "hello-world!"},
		},
		Started: true,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}

func TestContainerExecDebianSlim(t *testing.T) {
	ctx := context.Background()
	container, e := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			Image: Debian.AWS_ECR_URI + "/" + Debian.DOCKER_IMAGE_GROUP + "/" + Debian.DOCKER_IMAGE + ":" + Debian.DOCKER_IMAGE_TAG,
			Cmd:   []string{"echo", "hello-world!"},
		},
		Started: true,
	})
	require.NoError(t, e)
	testcontainers.CleanupContainer(t, container)
}
