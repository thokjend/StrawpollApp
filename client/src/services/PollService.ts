export const fetchPollData = async (pollId: string) => {
  try {
    const response = await fetch(`http://localhost:8080/poll/${pollId}`);
    const result = await response.json();

    if (response.ok) {
      return result;
    }
  } catch (error) {
    console.log("Error fetching poll");
  }
};
