export const fetchProtectedData = async () => {
  const token = localStorage.getItem("sessionToken");
  if (!token) {
    console.error("No session token found. Please log in.");
    return null;
  }

  try {
    const response = await fetch("http://localhost:8080/protected-endpoint", {
      method: "GET",
      headers: {
        Authorization: `Bearer ${token}`,
        "Content-Type": "application/json",
      },
    });

    if (response.ok) {
      return await response.json();
    } else {
      console.error("Unauthorized request. Token may have expired.");
      localStorage.removeItem("sessionToken");
      return null;
    }
  } catch (error) {
    console.error("Error fetching protected data:", error);
    return null;
  }
};
