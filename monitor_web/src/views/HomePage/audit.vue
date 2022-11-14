<template>
  <div class="home-page">
    <template v-if="token === ''">
      <el-input v-model="token" placeholder="请输入token"
                style="width: 200px;display: inline-block;margin-right: 10px"></el-input>
      <el-button type="primary" @click="login">登录</el-button>
    </template>
    <template v-else>
      <div class="table-title">
        <div class="table-title-left">
          状态：
          <el-select v-model="page.status" placeholder="请选择" @change="changeStatus" :clearable="true">
            <el-option
              v-for="item in statusFilter"
              :key="item.value"
              :label="item.name"
              :value="item.value">
            </el-option>
          </el-select>
          <el-button type="primary" @click="changeSell" :loading="loading" style="margin-left: 10px">{{
              btnName
            }}
          </el-button>
        </div>
        <div class="table-title-right">
          <el-button type="primary" @click="logout" style="margin-right: 10px">退出登录
          </el-button>
        </div>
      </div>

      <el-skeleton
        :loading="loading"
        animated
        :throttle="500"
      >
        <template slot="template">
          <div style="padding: 14px;">
            <el-skeleton-item variant="h1"/>
            <el-skeleton-item variant="h1"/>
          </div>
        </template>
        <template>
          <el-table :data="tableData" style="width: 100%">
            <el-table-column prop="Name" label="id"></el-table-column>
            <el-table-column prop="URL" label="链接"></el-table-column>
            <el-table-column prop="Price" label="状态"></el-table-column>
            <el-table-column prop="Stock" label="提交日期" width="120">
              <template slot-scope="scope">
                <el-tag
                  :type="scope.row.Stock ? 'success' : 'danger'"
                  disable-transitions>{{ scope.row.Stock ? '有货' : '缺货' }}
                </el-tag>
              </template>
            </el-table-column>
            <el-table-column
              label="操作"
              width="100">
              <template slot-scope="scope">
                <template v-if="scope.row.Stock">
                  <el-button @click="open_page(scope.row.URL)" type="success" size="small">购买</el-button>
                </template>
                <template v-else>
                  <el-button type="danger" size="small" disabled>购买</el-button>
                </template>
              </template>
            </el-table-column>
          </el-table>
          <el-pagination
            @size-change="handleSizeChange"
            @current-change="handleCurrentChange"
            :current-page.sync="currentPage2"
            :page-sizes="[20, 50, 100]"
            :page-size="page.size"
            layout="sizes, prev, pager, next"
            :total="total">
          </el-pagination>
        </template>
      </el-skeleton>
    </template>
  </div>
</template>
<script>
import { Audit, AuditList } from '@/apis/data.api';
import userStore from '@/store/modules/app';
import { useRoute } from '@/router';
import { open_page } from '@/utils/util';

export default {
  data() {
    return {
      loading: true,
      btnName: '加载中',
      subLoading: false,
      subBtnName: '确 定',
      tableData: [],
      store: userStore(),
      useRoute,
      open_page,
      token: '',
      id: undefined,
      total: 0,
      page: {
        page: 1,
        size: 10,
        status: undefined,
      },
      statusFilter: [
        {
          name: '未处理',
          value: 1,
        },
        {
          name: '已处理',
          value: 2,
        },
      ],
    };
  },
  mounted() {
    this.token = localStorage.getItem('secret');
    if (this.token) {
      this.getAuditList();
    }
  },
  computed: {},
  methods: {
    login() {
      localStorage.setItem('secret', this.token);
      this.getAuditList();
    },
    logout() {
      this.token = '';
      localStorage.setItem('secret', this.token);
    },
    async getAuditList() {
      let res = await AuditList();
      if (res.code === 403) {
        this.logout();
        return;
      }
      this.tableData = res.data.data;
      this.total = res.data.total;
    },
    async Audit() {
      let res = await Audit();
      this.sellers = res.data;
    },

    async getPlistApi() {
      this.loading = true;
      this.tableData = [];
      let res = await GetPlistApi({ id: this.id, stock: this.status });
      this.tableData = res.data;
      this.btnName = '刷新';
      this.loading = false;
    },

    changeStatus() {
      this.getAuditList(this.page);
    },
    submit() {
      this.subLoading = true;
      this.subBtnName = '提交中';
      this.$refs.ruleForm.validate(async (valid) => {
        if (!valid) {
          console.log('error submit!!');
          this.subLoading = false;
          this.subBtnName = '确 定';
          return false;
        }
        let res = await SubmitApi(this.form);
        this.form = {
          name: 'a',
          url: '',
          price: '',
          productName: '',
        };
        this.subLoading = false;
        this.subBtnName = '确 定';
        this.dialogFormVisible = false;
      });
    },
  },
  components: {},
};
</script>
<style lang="scss" scoped>
.home-page {
  .test-text {
    color: v-bind(textColor);
  }

  .table-title {
    display: flex;
    margin: 20px 0;

    .table-title-right {
      margin-left: auto;
    }
  }
}
</style>
