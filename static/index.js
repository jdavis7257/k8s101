var app = new Vue({
  el: "#app",
  data: {
    comments: [],
    newCommentText: "test"
  },
  mounted() {
    axios.get("api/v1/comment/getAll").then(response => {
      this.comments = response.data;
      console.log(response.data);
    });
  },
  methods: {
    postData() {
      var bodyFormData = new FormData();
      bodyFormData.set("comment", this.newCommentText);
      console.log(this.newCommentText);
      console.log("Posting data");
      axios({
        method: "POST",
        url: "api/v1/comment/new",
        data: bodyFormData,
        config: {
          headers: { "Content-Type": "application/x-www-form-urlencoded" }
        }
      })
        .then(function(response) {
          //handle success
          console.log(response);
          this.newCommentText = "";
        })
        .catch(function(response) {
          //handle error
          console.log(response);
        });
      return false;
    }
  }
});
