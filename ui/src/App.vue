<template>
    <div>
    <el-table :data="filteredUsers" style="width: 85%;margin: auto">
    <el-table-column prop="id" label="ID" width="60" />
    <el-table-column prop="name" label="Name" width="120" />
    <el-table-column prop="email" label="Email" width="180" />
    <el-table-column prop="phone" label="Phone" width="180" />
    <el-table-column align="right">
      <template #header>
        <el-input v-model="search" size="small" placeholder="Type to search" />
      </template>
      <template #default="scope">
        <el-button size="small" @click="handleEdit(scope.$index, scope.row)"
          >Edit</el-button
        >
        <el-button
          size="small"
          type="danger"
          @click="handleDelete(scope.$index, scope.row)"
          >Delete</el-button
        >
      </template>
    </el-table-column>
  </el-table>
  </div>
</template>

<script>


export default {
  name: 'App',
  components: {
    
  },
  data() {
    return {
      search : '',
      users : []
    }
  },
  created() {
    this.getUsers();
  },
  computed: {
    filteredUsers() {
      var self = this;
      return this.users.filter(item => {
        return self.search == '' || item.name.toLowerCase().includes(self.search.toLowerCase())
      });
    }
  },
  methods: {
    getUsers() {
      fetch('http://localhost:8086/api/v1/users/')
        .then(response => response.json())
        .then(json => {
          this.users = json.data;
          console.log(this.users );
        })
    }
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
