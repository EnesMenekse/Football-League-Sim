document.getElementById("simulateButton").addEventListener("click", simulate);
document.getElementById("nextWeekButton").addEventListener("click", nextWeek);
window.addEventListener("load", loadStandings);
window.addEventListener("load", loadMatches);

function simulate() {
  fetch("/simulate")
    .then(handleResponse)
    .then(() => {
      loadStandings();
      loadMatches();
      alert("Simulation Completed!");
    })
    .catch(handleError);
}

function nextWeek() {
  fetch("/next-week")
    .then(handleResponse)
    .then(() => {
      loadStandings();
      loadMatches();
      alert("Next Week Processed!");
    })
    .catch(handleError);
}

function loadStandings() {
  fetch("/standings")
    .then(handleResponse)
    .then(renderStandings)
    .catch(handleError);
}

function loadMatches() {
  fetch("/matches").then(handleResponse).then(renderMatches).catch(handleError);
}

function handleResponse(response) {
  if (!response.ok) {
    throw Error(response.statusText);
  }
  return response.json();
}

function renderStandings(data) {
  const tbody = document.querySelector("#standingsTable tbody");
  tbody.innerHTML = "";
  data.forEach((team) => {
    const row = document.createElement("tr");
    row.innerHTML = `
      <td>${team.Name}</td>
      <td>${team.Points}</td>
      <td>${team.Wins}</td>
      <td>${team.Draws}</td>
      <td>${team.Losses}</td>
      <td>${team.GoalsFor}</td>
      <td>${team.GoalsAgainst}</td>
    `;
    tbody.appendChild(row);
  });
}

function renderMatches(data) {
  const tbody = document.querySelector("#matchesTable tbody");
  tbody.innerHTML = "";
  data.forEach((match) => {
    const row = document.createElement("tr");
    row.innerHTML = `
      <td>${match.home_team}</td>
      <td>${match.away_team}</td>
      <td>${match.home_goals}</td>
      <td>${match.away_goals}</td>
    `;
    tbody.appendChild(row);
  });
}

function handleError(error) {
  console.error("Error:", error);
}
