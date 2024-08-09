import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";

export default function Home() {
  return (
    <div className="h-screen flex justify-center items-center">
      <form className="w-full max-w-md flex flex-col gap-4">
        <h1 className="text-center text-2xl">晓智科技</h1>
        <Input type="text" placeholder="请输入用户名" />
        <Textarea placeholder="请输入" />
        <Button type="submit">登录</Button>
      </form>
    </div>
  );
}
