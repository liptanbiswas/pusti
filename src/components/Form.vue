<template>
    <div class="row justify-content-center">
      <div>
        <h1 v-if="!status">Please provide your feedback!</h1>

        <b-form v-if="!status" @submit="sendForm">
          <b-form-group>
            <b-form-input type="text" name="name" v-model="name" placeholder="Enter your name..."></b-form-input>
          </b-form-group>
          <b-form-group>
            <b-form-input type="email" name="email" v-model="email" placeholder="Enter your email..." required>
            </b-form-input>
          </b-form-group>
          <b-form-group>
            <b-form-textarea type="text" name="message" v-model="message" placeholder="Leave a feedback..." required>
            </b-form-textarea>
          </b-form-group>
          <b-form-group>
            <b-button block variant="primary" type="submit">Send</b-button>
          </b-form-group>
        </b-form>

        <h1 v-if="status === 'success'">Thank you, your feedback is submitted!</h1>
        <h1 v-if="status === 'error'">Oops, something went wrong. Please try again.</h1>
      </div>
    </div>
  </template>

  <script>
    export default {
      name: 'Form',
      data: function () {
        return {
          status: null,
          name: null,
          email: null,
          message: null
        }
      },
      methods: {
        sendForm: function (event) {
          event.preventDefault()
          fetch('/api/sendmail', {
            method: 'POST',
            headers: {
              'Content-Type': 'application/json',
              Accept: 'application/json'
            },
            body: JSON.stringify({ name: this.name, email: this.email, message: this.message })
          })
            .then((response) => {
              if (response.status < 299) {
                this.status = 'success'
              }
              else {
                this.status = 'error'
              }
            })
        }
      }
    }
  </script>
