export async function apiRequest(url, options = {}) {
  const accessToken = localStorage.getItem("jwt");

  // Add access token to header
  const headers = {
    ...(options.headers || {}),
    Authorization: `Bearer ${accessToken}`,
  };

  let res = await fetch(url, { ...options, headers });

  // If unauthorized → try refresh token
  if (res.status === 401) {
    const refreshToken = localStorage.getItem("refresh_token");

    if (!refreshToken) {
      return redirectToLogin();
    }

    const refreshRes = await fetch("http://127.0.0.1:8080/refresh", {
      method: "POST",
      headers: {
        refresh_token: refreshToken,
      },
    });

    // If refresh also fails → login again
    if (!refreshRes.ok) {
      return redirectToLogin();
    }

    // Store new access token
    const data = await refreshRes.json();
    localStorage.setItem("jwt", data.access_token);

    // Retry original request
    const retryHeaders = {
      ...(options.headers || {}),
      Authorization: `Bearer ${data.access_token}`,
    };

    res = await fetch(url, { ...options, headers: retryHeaders });
  }

  return res;
}

function redirectToLogin() {
  localStorage.removeItem("jwt");
  localStorage.removeItem("refresh_token");
  window.location.href = "/login";
}
