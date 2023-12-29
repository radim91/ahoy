import { useEffect, useState } from 'react';
import Layout from './components/Layout';

function App() {
    const [data, setData] = useState({});

    useEffect(() => {
        fetch('http://localhost:3333/api/projects')
            .then(response => response.json())
            .then(data => {
                setData(data);
            })
        ;

        /* fetch('http://localhost:3333/api/containers') */
        /*     .then(response => response.json()) */
        /*     .then(data => { */
        /*         setContainers(data); */
        /*     }) */
        /* ; */
        /**/
        /* fetch('http://localhost:3333/api/images') */
        /*     .then(response => response.json()) */
        /*     .then(data => { */
        /*         setImages(data); */
        /*     }) */
        /* ; */
        /**/
        /* fetch('http://localhost:3333/api/networks') */
        /*     .then(response => response.json()) */
        /*     .then(data => { */
        /*         setNetworks(data); */
        /*     }) */
        /* ; */
    }, []);

    return (
        <Layout>
            <div>
                <h3>Projects</h3>
                <ol>
                {Object.values(data).map(item => (
                    <li key={item.Hash}>{item.Name}</li>
                ))}
                </ol>
                {/* <h3>Containers</h3> */}
                {/* <ol> */}
                {/* {Object.values(containers).map(container => ( */}
                {/*     <li key={container.Id}>{container.Names}</li> */}
                {/* ))} */}
                {/* </ol> */}
                {/* <h3>Images</h3> */}
                {/* <ol> */}
                {/* {Object.values(images).map(image => ( */}
                {/*     <li key={image.Id}>{image.RepoTags}</li> */}
                {/* ))} */}
                {/* </ol> */}
                {/* <h3>Networks</h3> */}
                {/* <ol> */}
                {/* {Object.values(networks).map(network => ( */}
                {/*     <li key={network.Id}>{network.Name}</li> */}
                {/* ))} */}
                {/* </ol> */}
            </div>
        </Layout>
    );
}

export default App;
