import router from "../router";

export const ValidateToken = async () => {
  const token = localStorage.getItem("sessionToken");

  if (!token) {
    router.push({ name: "Login" });
    return false;
  }

  try {
    const response = await fetch("http://localhost:8080/protected-route", {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        Authorization: token,
      },
    });

    if (!response.ok) {
      if (response.status === 401) {
        localStorage.removeItem("sessionToken"); // Remove invalid token
        router.push({ name: "Login" }); // Redirect to login
      }
      throw new Error("Unauthorized access");
    }

    const result = await response.json();
    //console.log("Session verified:", result);
    return result;
  } catch (error) {
    console.error("Error verifying session", error);
    localStorage.removeItem("sessionToken");
    router.push({ name: "Login" });
    return false;
  }
};
