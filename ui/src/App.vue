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
        <el-popconfirm
          confirm-button-text="Yes"
          cancel-button-text="No, Thanks"
          :icon="InfoFilled"
          icon-color="#626AEF"
          title="Are you sure to delete this?"
          @confirm="confirmDeleteEvent"
          @cancel="cancelEvent"
        >
          <template #reference>
            <el-button
                size="small"
                type="danger"
                @click="handleDelete(scope.$index, scope.row)"
                >Delete</el-button>
          </template>
        </el-popconfirm>
        
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

import { ElNotification } from 'element-plus'
export default {
  name: 'App',
  components: {
    
  },
  data() {
    return {
      search : '',
      users : [],
      selected_user : {},
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
    deleteUsers(id) {
      fetch('http://localhost:8086/api/v1/users/'+id, {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json'
        }
      })
        .then(response => response.json())
        .then(json => {
          console.log(json);
          this.getUsers();
          ElNotification({
            title: 'Success',
            message: 'User deleted successfully',
            type: 'success',
          })
        })
    },
    prevPage(val) {
      this.pagination.page = val;
    },
    nextPage(val) {
      this.pagination.page = val;
    },
    handleCurrentChange(val) {
      this.pagination.page = val;
      this.getUsers();
    },
    handleDelete(index, row){
      console.log(index, row.id);
      this.selected_user = row;
    },
    confirmDeleteEvent() {
      console.log(this.selected_user)
      this.deleteUsers(this.selected_user.id);
    },    
    cancelEvent() {
      console.log('cancel');
      this.selected_user = {};
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
