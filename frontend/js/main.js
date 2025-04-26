const touchedInputs = new Set();

///////////////////////////////////////////
async function fetchArtists() {
  try {
    const res = await fetch("http://localhost:8080/api/artists");
    if (!res.ok) {
      // If backend returned an error with JSON body
      const errorBody = await res.json();
      console.log(errorBody);

      throw {
        status: res.status,
        message: errorBody.message || "Server error",
        details: errorBody.details || "",
      };
    }
    // At this point, we are sure it's a valid response
    const data = await res.json();
    renderArtists(data);
  } catch (err) {
    renderError({
      status: err.status || "Network Error",
      message: err.message || "Something went wrong",
      details: err.details || "",
    });
    // return [];
  }
}
function renderArtists(artists) {
  const container = document.getElementById("container");
  container.innerHTML = "";

  if (!artists || artists.length === 0) {
    container.innerHTML = "No Artist Found";
    return;
  }

  artists.forEach((artist) => {
    const card = document.createElement("div");
    card.className = "card";

    card.innerHTML = `
      <img src="${artist.image}" alt="${artist.name}">
      <h2>${artist.name}</h2>
      <button class="details-btn" onclick="artistDetails(${artist.id})">Details</button>
    `;

    container.appendChild(card);
  });
}
function updateRangeValues() {
  const ranges = document.querySelectorAll('input[type="range"]');

  ranges.forEach((range) => {
    const label = document.getElementById(`${range.id}-value`);
    if (!label) return;

    // Initial value display
    label.textContent = range.value;

    // Update display + mark as touched
    range.addEventListener("input", () => {
      label.textContent = range.value;
      touchedInputs.add(range.name);
    });
  });
}
function renderError(err) {
  document.body.innerHTML = "";

  const errorContent = document.createElement("div");
  errorContent.classList.add("error-box");

  errorContent.innerHTML = `
    <h1>Oooops 😬</h1>
    <p><strong>${err.status || "Error"}</strong> | ${
    err.message || "Something went wrong."
  }</p>
    ${err.details ? `<pre>${err.details}</pre>` : ""}
    <button class="backhome-btn"><a href="/frontend/">Back Home</a></button>
  `;

  document.body.appendChild(errorContent);
}
updateRangeValues();
fetchArtists();
///////////////////////////////////////////////////

async function artistDetails(id) {
  try {
    const res = await fetch(
      `http://localhost:8080/api/artists/details?id=${id}`
    );
    const data = await res.json(); // always parse JSON

    if (!res.ok) {
      const error = new Error(res.statusText);
      error.status = res.status;
      throw error;
    }

    displayArtistDetails(data);
  } catch (err) {
    // network failure or invalid JSON
    renderError(err);
  }
  // });
  // });
}

function displayArtistDetails(data) {
  const container = document.getElementById("container");
  const panel = document.getElementById("artist-details-panel");

  const { ArtistInfo, Locations, Dates, Relations } = data;

  const locations = Locations.locations
    .map((loc) => `<li>${loc}</li>`)
    .join("");
  const concertDates = Dates.dates.map((date) => `<li>${date}</li>`).join("");
  const members = ArtistInfo.members.map((m) => `<li>${m}</li>`).join("");

  const datesByLocation = Object.entries(Relations.datesLocations)
    .map(([place, dates]) => {
      const prettyPlace = place;
      const dateList = dates.map((d) => `<li>${d}</li>`).join("");
      return `<div><h4>${prettyPlace}</h4><ul>${dateList}</ul></div>`;
    })
    .join("");
  container.innerHTML = "";
  container.innerHTML = `
      <div class="card">
        <img src="${ArtistInfo.image}" alt="${ArtistInfo.name}" />
        <div>
          <h2>${ArtistInfo.name}</h2>
          <p>First Album: ${ArtistInfo.firstAlbum}</p>
          <p>Creation Year: ${ArtistInfo.creationDate}</p>
        </div>
      </div>
  
      <div class="card">
        <h3>Members</h3>
        <ul>${members}</ul>
      </div>
  
      <div class="card">
        <h3>Locations</h3>
        <ul>${locations}</ul>
      </div>
  
      <div class="card">
        <h3>Concert Dates</h3>
        <ul>${concertDates}</ul>
      </div>
  
      <div class="card">
        <h3>Concerts by Location</h3>
        ${datesByLocation}
      </div>
      <button class="backhome-btn"><a href="/frontend/">Back Home</a></button>
    `;
}

//////////////////////////////////////////////////////////

function setupFilters() {
  const form = document.getElementById("filter-form");

  form.addEventListener("submit", async (e) => {
    e.preventDefault();

    const filters = {
      creationDateFrom: touchedInputs.has("creation-from")
        ? form.querySelector('input[name="creation-from"]').value
        : "0",

      creationDateTo: touchedInputs.has("creation-to")
        ? form.querySelector('input[name="creation-to"]').value
        : "0",

      firstAlbumFrom: touchedInputs.has("album-from")
        ? form.querySelector('input[name="album-from"]').value
        : "0",

      firstAlbumTo: touchedInputs.has("album-to")
        ? form.querySelector('input[name="album-to"]').value
        : "0",

      members:
        form.querySelector('input[name="members"]:checked')?.value || "0",

      concertDates: form.querySelector("#location-input")?.value.trim() || "",
    };

    try {
      const res = await fetch("http://localhost:8080/api/artists/filter", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(filters),
      });
      const filteredArtists = await res.json();
      if (!res.ok) {
        const error = new Error(res.statusText);
        error.status = res.status;
        throw error;
      }
      renderArtists(filteredArtists);
    } catch (err) {
      // network failure or invalid JSON
      renderError(err);
    }
  });
}
setupFilters();
///////////////////////////////////////////////////////

async function search() {
    
    document.addEventListener("DOMContentLoaded", () => {
        const searchInput = document.getElementById("search");
        const suggestionsBox = document.getElementById("suggestions");
        const container = document.getElementById("container");
      
        searchInput.addEventListener("input", async (event) => {
          const query = event.target.value.trim();
          console.log(query);
          
      
          if (query.length > 0) {
            try {
              const response = await fetch(`http://localhost:8080/api/artists/search?find=${query}`);
              const { Artists , Sugestions} = await response.json();
              console.log(Sugestions , 'Artiste');
              
      
              renderSuggestions(Sugestions);
              renderArtists(Artists);
            } catch (error) {
              console.error("API error:", error);
              suggestionsBox.style.display = "none";
              container.innerHTML = "<p>Error loading artists.</p>";
            }
          } else {
            suggestionsBox.style.display = "none";
            container.innerHTML = "";
          }
        });
      
      });
      
}

function renderSuggestions(suggestions) {
    const suggestionsBox = document.getElementById("suggestions");
    if (!suggestions || suggestions.length === 0) {
      suggestionsBox.style.display = "none";
      return;
    }

    suggestionsBox.innerHTML = `<ul>${suggestions
      .map(item => `<li>${item}</li>`)
      .join("")}</ul>`;
    suggestionsBox.style.display = "block";

    suggestionsBox.querySelectorAll("li").forEach(li => {
      li.addEventListener("click", () => {
        searchInput.value = li.textContent;
        suggestionsBox.style.display = "none";
        searchInput.dispatchEvent(new Event("input")); // Trigger new search
      });
    });
  }



search();

///////////////////////////////////////////////////////

function darckMode(params) {
  const themeButton = document.getElementById("theme");
  themeButton.addEventListener("click", () => {
    document.body.classList.toggle("dark-mode");
    const isDark = document.body.classList.contains("dark-mode");
    themeButton.textContent = isDark ? "🌞" : "🌙"; // Tema butonunun ikonunu değiştir
  });
}
document.addEventListener("DOMContentLoaded", () => {
  const toggleBtn = document.getElementById("toggle-filters");
  const filterSection = document.getElementById("filters");

  toggleBtn.addEventListener("click", () => {
    filterSection.classList.toggle("hidden");
  });
});

darckMode();
