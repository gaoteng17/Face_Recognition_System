package gaoteng17.testpost;

import android.os.Bundle;
import java.io.InputStream;
import java.io.OutputStream;
import java.io.PrintWriter;
import java.net.HttpURLConnection;
import java.net.URL;

import android.app.Activity;
import android.os.Handler;
import android.os.Message;
import android.view.View;
import android.view.View.OnClickListener;
import android.widget.Button;
import android.widget.TextView;
import android.widget.Toast;

public class MainActivity extends Activity {
    public Button loginBtn;
    public TextView loginUserName;
    public TextView loginSchoolId;
    public TextView loginClass;
    public static String API="http://123.56.223.156/";
    public LoginHandler loginHandler;
    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        //获取View对象
        loginBtn=(Button) findViewById(R.id.loginBtn);
        loginUserName=(TextView) findViewById(R.id.loginUsername);
        loginSchoolId=(TextView) findViewById(R.id.loginSchoolId);
        loginClass=(TextView) findViewById(R.id.loginClass);
        //给View对象设置点击事件
        loginBtn.setOnClickListener(new OnClickListener() {
            @Override
            public void onClick(View arg0) {
                //开启新线程
                Thread loginThread=new Thread(new LoginRunable());
                loginThread.start();
            }
        });
        loginHandler=new LoginHandler();
    }
    //实现Runable接口,开启新线程
    class LoginRunable implements Runnable{
        @Override
        public void run() {
            try {
                URL url=new URL(API);
                HttpURLConnection http=(HttpURLConnection) url.openConnection();
                http.setRequestMethod("POST");
                http.setDoInput(true);
                http.setDoOutput(true);
                OutputStream ops=http.getOutputStream();
                PrintWriter pw=new PrintWriter(ops);
                String Username=loginUserName.getText().toString();
                String Class=loginClass.getText().toString();
                String SchoolId=loginSchoolId.getText().toString();
                //String s = new String("{\"school_id\": \"20164762\", \"name\": \"jieli\", \"class\": \"1603\", \"options\": \"delete\", \"parameters\": \"\"}");
                pw.write("{\"school_id\": \""+SchoolId+"\", \"name\": \""+Username+"\", \"class\": \""+Class+"\", \"options\": \"delete\", \"parameters\": \"\"}");
                //pw.write(s);
                pw.flush();

                InputStream ins=http.getInputStream();
                byte[] buffer = new byte[1024];
                int length=0;
                StringBuilder sb=new StringBuilder();
                while((length=ins.read(buffer))!=-1){
                    sb.append(new String(buffer,0,length));
                }

                Message msg=new Message();
                msg.what=1;
                msg.obj=sb.toString();
                loginHandler.sendMessage(msg);
            } catch (Exception e) {
                // TODO Auto-generated catch block
                e.printStackTrace();
            }

        }
    }
    //传递消息的handle
    class LoginHandler extends Handler{
        @Override
        public void handleMessage(Message msg) {
            String loginResponse=(String) msg.obj;
            //System.out.println(loginResponse);
            Toast.makeText(MainActivity.this, loginResponse, 10).show();
        }
    }
}