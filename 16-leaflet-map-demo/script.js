const map=L.map("map").setView([25.4358,81.8463],15);

L.tileLayer(
    "https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png",
    {
        attribution:"© OpenStreetMap contributors"
    }
).addTo(map);

const hazards=[
    {
        type:"Pothole",
        severity:"High",
        latitude:25.4358,
        longitude:81.8463
    },
    {
        type:"Road Crack",
        severity:"Medium",
        latitude:25.4380,
        longitude:81.8500
    },
    {
        type:"Damaged Road",
        severity:"Low",
        latitude:25.4330,
        longitude:81.8420
    }
];

hazards.forEach(hazard=>{
    const marker=L.marker([
        hazard.latitude,
        hazard.longitude
    ]).addTo(map);

    marker.bindPopup(`
        <b>${hazard.type}</b><br>
        Severity: ${hazard.severity}
    `);
});