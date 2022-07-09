<template>
    <div>
    <el-table :data="users" style="width: 85%;margin: auto">
    <el-table-column prop="id" label="ID" width="60" />
    <el-table-column prop="name" label="Name" width="120" />
    <el-table-column prop="email" label="Email" width="180" />
    <el-table-column prop="phone" label="Phone" width="180" />
    <el-table-column align="right">
      <template #header>
        <el-input v-debounce:300ms="getUsers" v-model="search" size="small" placeholder="Type to search" />
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
  <el-pagination background layout="prev, pager, next" 
    :total="pagination.total"  
    @prev-click="prevPage" 
    @next-click="nextPage" 
    @current-change="handleCurrentChange"
    :current-page="pagination.page"
    :page-size="pagination.limit"  
  />
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
      users : [],
      pagination : {
        page : 1,
        limit : 12,
        total : 0
      },
    }
  },
  created() {
    this.getUsers();
  },
  methods: {
    getUsers() {
      fetch('http://localhost:8086/api/v1/users/?search='+this.search+'&page='+this.pagination.page+'&limit='+this.pagination.limit)
        .then(response => response.json())
        .then(json => {
          this.users = json.data;
          this.pagination = json.pagination;
        })
    },
    prevPage(val) {
      this.pagination.page = val
    },
    nextPage(val) {
      this.pagination.page = val;
    },
    handleCurrentChange(val) {
      this.pagination.page = val;
      this.getUsers();
    },
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
