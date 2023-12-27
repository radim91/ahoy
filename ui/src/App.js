import { useEffect, useState } from 'react';

function App() {
    const [projects, setProjects] = useState([]);
    const [containers, setContainers] = useState([]);
    const [images, setImages] = useState([]);
    const [networks, setNetworks] = useState([]);

    useEffect(() => {
        fetch('http://localhost:3333/api/projects')
            .then(response => response.json())
            .then(data => {
                setProjects(data);
            })
        ;

        fetch('http://localhost:3333/api/containers')
            .then(response => response.json())
            .then(data => {
                setContainers(data);
            })
        ;

        fetch('http://localhost:3333/api/images')
            .then(response => response.json())
            .then(data => {
                setImages(data);
            })
        ;

        fetch('http://localhost:3333/api/networks')
            .then(response => response.json())
            .then(data => {
                setNetworks(data);
            })
        ;
    }, []);

    return (
    <div>
        <h3>Projects</h3>
        <ol>
        {Object.values(projects).map(project => (
            <li key={project.Hash}>{project.Name}</li>
        ))}
        </ol>
        <h3>Containers</h3>
        <ol>
        {Object.values(containers).map(container => (
            <li key={container.Id}>{container.Names}</li>
        ))}
        </ol>
        <h3>Images</h3>
        <ol>
        {Object.values(images).map(image => (
            <li key={image.Id}>{image.RepoTags}</li>
        ))}
        </ol>
        <h3>Networks</h3>
        <ol>
        {Object.values(networks).map(network => (
            <li key={network.Id}>{network.Name}</li>
        ))}
        </ol>
    </div>
    );
}

export default App;
