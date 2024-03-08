package entity

import (
	"os/exec"
	"strings"
)

type Project struct {
	Hash     string
	Name     string
	File     string
	Workdir  string
	Services []Service
}

type Service struct {
	Name  string
	Hash  string
	Image string
}

func GetProjects() map[string]Project {
	containers := GetContainers()
	projects := make(map[string]Project, 0)

	for _, container := range containers {
		if container.Labels["com.docker.compose.project"] != "" {
			projectName := container.Labels["com.docker.compose.project"]
			existingProject, ok := projects[projectName]
			service := Service{
				Name:  container.Labels["com.docker.compose.service"],
				Hash:  container.Labels["com.docker.compose.config-hash"],
				Image: container.Labels["com.docker.compose.image"],
			}

			var project Project = Project{}

			if ok {
				existingProject.Services = append(existingProject.Services, service)
				project = existingProject
			} else {
				project = Project{
					Hash:    container.Labels["com.docker.compose.config-hash"],
					Name:    projectName,
					File:    container.Labels["com.docker.compose.project.config_files"],
					Workdir: container.Labels["com.docker.compose.project.working_dir"],
					Services: []Service{
						service,
					},
				}
			}

			projects[project.Name] = project
		}
	}

	return projects
}

func GetProject(name string) Project {
	return GetProjects()[name]
}

func StartProject(project Project) {
	args := setArgs(project)
	args = append(args, "up", "-d")

	cmd := exec.Command("docker-compose", args...)
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func StopProject(project Project) {
	args := setArgs(project)
	args = append(args, "stop")

	cmd := exec.Command("docker-compose", args...)
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func RestartProject(project Project) {
	args := setArgs(project)
	args = append(args, "restart")

	cmd := exec.Command("docker-compose", args...)
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func GetProjectStatus(projectName string) string {
	project := GetProject(projectName)

	args := setArgs(project)
	args = append(args, "ps", "--format", "json")

	output, err := exec.Command("docker-compose", args...).Output()

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(output), "\n")
	jsonObjects := strings.Join(lines[0:len(lines)-1], ",")

	return "[" + jsonObjects + "]"
}

func DownProject(project Project) {
	args := setArgs(project)
	args = append(args, "down")

	cmd := exec.Command("docker-compose", args...)
	err := cmd.Run()

	if err != nil {
		panic(err)
	}
}

func setArgs(project Project) []string {
	args := []string{}

	files := strings.Split(project.File, ",")
	for _, file := range files {
		args = append(args, "-f")
		args = append(args, file)
	}

	return args
}
