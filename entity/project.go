package entity

type Project struct {
    Hash string
    Name string
    File string
    Workdir string
    Services []Service
}

type Service struct {
    Name string
    Hash string
    Image string
}

func GetProjects() (map[string]Project) {
    containers := GetContainers()
    projects := make(map[string]Project, 0)

    for _, container := range containers {
        if container.Labels["com.docker.compose.project"] != "" {
            projectName := container.Labels["com.docker.compose.project"]
            existingProject, ok := projects[projectName]
            service := Service{
                Name: container.Labels["com.docker.compose.service"],
                Hash: container.Labels["com.docker.compose.config-hash"],
                Image: container.Labels["com.docker.compose.image"],
            }

            var project Project = Project{}

            if ok {
                existingProject.Services = append(existingProject.Services, service)
                project = existingProject
            } else {
                project = Project{
                    Hash: container.Labels["com.docker.compose.config-hash"],
                    Name: projectName,
                    File: container.Labels["com.docker.compose.project.config_files"],
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

func GetProject(name string) (Project) {
    return GetProjects()[name]
}
